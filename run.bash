#!/bin/bash
curl "http://localhost:9999/Tom" &
curl "http://localhost:9999/Sam" &
curl "http://localhost:9999/Tom" &
curl "http://localhost:9999/Jack" &
curl "http://localhost:9999/Sam" &
wait