#!/bin/bash

TZ=US/Eastern ./clock3 -port 8001 &
TZ=Asia/Tokyo ./clock3 -port 8002 &
TZ=Europe/London ./clock3 -port 8003 &
