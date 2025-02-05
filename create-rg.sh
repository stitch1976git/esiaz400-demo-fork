#!/bin/bash

read "Enter Location: "  loc
read "Enter Name: " name

az group create -n $name -l $loc