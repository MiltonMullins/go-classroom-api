#!/bin/sh

# Inicia el productor en segundo plano
./producer &

# Inicia el consumidor en primer plano
./consumer