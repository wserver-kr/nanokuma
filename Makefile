# SPDX-License-Identifier: GPL-2.0-or-later
# Copyright (C) 2022-2026 WSERVER

SERVER_DIR = server/
TARGET = nanokuma

all: server

server: $(SERVER_DIR)
	make -C $(SERVER_DIR)

clean:
	@rm -f $(TARGET)

.PHONY: all server clean
