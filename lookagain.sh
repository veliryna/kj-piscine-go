#!/bin/bash
find ./ -type f -name "*.sh" -printf "%f\n" | sort -r | sed 's/.sh//g'