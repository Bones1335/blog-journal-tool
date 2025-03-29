package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type MetaData struct {
	Title      string    `yaml:"title"`
	Date       time.Time `yaml:"date"`
	URL        string    `yaml:"url"`
	Categories []string  `yaml:"categories"`
	Tags       []string  `yaml:"tags"`
	Content    string    `yaml:"-"`
}

func handlerNewFile(s *state, cmd command) error {
	editor := s.config.Editor
	tmpFile, err := os.CreateTemp("", "index*.md")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	templateText := fmt.Sprintf(`---
title: 'Testing File'
date: %v
url: 
categories: 
    - example 1
    - example 2
tags:
    - example 3
    - example 4
---
`, time.Now().Format(time.RFC3339))

	file, err := os.Create(tmpFile.Name())
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(templateText))
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		file.Close()
		return err
	}

	file.Close()

	open := exec.Command(editor, tmpFile.Name())
	open.Stdin = os.Stdin
	open.Stdout = os.Stdout
	open.Stderr = os.Stderr

	err = open.Run()
	if err != nil {
		return err
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return err
	}

	finalFile := "index.md"

	file, err = os.Create(finalFile)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		file.Close()
		return nil
	}

	err = file.Sync()
	if err != nil {
		file.Close()
		return err
	}

	file.Close()

	savedContent, err := os.ReadFile(finalFile)

	data, metaData, err := parseContent(savedContent)

	if err := os.WriteFile(finalFile, data, 0644); err != nil {
		return nil
	}

	err = mvFile(finalFile, metaData.URL, s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func parseContent(content []byte) ([]byte, MetaData, error) {
	contentStr := string(content)

	re := regexp.MustCompile(`(?s)^---\n(.*?)\n---\n(.*)`)
	matches := re.FindStringSubmatch(contentStr)

	var metaData MetaData

	if len(matches) == 3 {
		err := yaml.Unmarshal([]byte(matches[1]), &metaData)
		if err != nil {
			fmt.Println("ERROR: Failed to parse YAML:", err)
			return []byte{}, MetaData{}, err
		}
		metaData.Content = matches[2]
	} else {
		metaData.Content = contentStr
	}

	metaData.URL = fmt.Sprintf("/%v", strings.ToLower(strings.Join(strings.Split(metaData.Title, " "), "-")))

	data, err := yaml.Marshal(metaData)
	if err != nil {
		return []byte{}, MetaData{}, err
	}

	formattedContent := fmt.Sprintf("---\n%s---\n%s", data, metaData.Content)
	return []byte(formattedContent), metaData, nil
}

func mvFile(file string, dest string, s *state, cmd command) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	var saveLocation string
	if cmd.Args[0] == "blog" {
		saveLocation = home + s.config.Blog
	} else if cmd.Args[0] == "journal" {
		saveLocation = home + s.config.Journal
	} else {
		return fmt.Errorf("error reading save location from command line input")
	}

	sourcePath := filepath.Join(currentDir, file)
	destDir := filepath.Join(saveLocation, dest)
	destPath := filepath.Join(destDir, file)

	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Rename(sourcePath, destPath)
	if err != nil {
		return err
	}

	fmt.Println("File moved successfully")

	return nil
}
