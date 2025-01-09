#!/bin/bash
export ISOLATE=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo "$ISOLATE"
INT=$(cat interviews/interview-"$ISOLATE")
echo "$INT"
echo "$MAIN_SUSPECT"
