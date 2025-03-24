package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func handlerNewFile(s *state, cmd command) error {
	editor := s.config.Editor
	tmpFile, err := os.CreateTemp("", "index.md")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	templateText := fmt.Sprintf(`---
title: 
url:
date: %v
categories: 
tags:
  - 
  - 
---
`, time.Now())

	if err := os.WriteFile(tmpFile.Name(), []byte(templateText), 0644); err != nil {
		return err
	}

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
	if err := os.WriteFile(finalFile, content, 0644); err != nil {
		return err
	}

	return nil
}
