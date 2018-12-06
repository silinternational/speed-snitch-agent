FROM golang:latest

# Install Deps
RUN apt-get update && apt-get install -y ruby-dev build-essential

# Cleaning
RUN apt-get clean

RUN gem install fpm