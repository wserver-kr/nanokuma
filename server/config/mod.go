// SPDX-License-Identifier: GPL-2.0-or-later
/*
 * nanokuma
 * Copyright (C) 2022-2026 WSERVER
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 */

package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type SSLConfig struct {
	Enable   bool   `toml:"enable"`
	KeyFile  string `toml:"key_file"`
	CertFile string `toml:"cert_file"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Prefix   string `toml:"prefix"`
}

type RawConfig struct {
	Host     string         `toml:"host"`
	Port     int            `toml:"port"`
	SSL      SSLConfig      `toml:"ssl"`
	Database DatabaseConfig `toml:"database"`
}

const CONFIG_PATH = "config.toml"

var (
	Get *RawConfig

	DefaultConfig RawConfig = RawConfig{
		Host: "localhost",
		Port: 8080,
		SSL: SSLConfig{
			Enable: false,
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     3306,
			Name:     "nanokuma",
			Username: "root",
			Password: "",
			Prefix:   "nk_",
		},
	}
)

func Load() error {
	var err error
	var buf []byte
	var raw []byte
	var data RawConfig

	buf, err = os.ReadFile(CONFIG_PATH)
	if err != nil {
		fmt.Printf("[nanokuma]: config.toml is not exists, creating new config file...\n")
		raw, err = toml.Marshal(&DefaultConfig)
		if err != nil {
			return err
		}

		err = os.WriteFile(CONFIG_PATH, raw, 0600)
		if err != nil {
			return err
		}
	}

	err = toml.Unmarshal(buf, &data)
	if err != nil {
		return err
	}

	Get = &data

	return nil
}

func Unload() error {
	if Get != nil {
		Get = nil
	}

	return nil
}
