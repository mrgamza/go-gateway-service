package config

import (
    "io/ioutil"
    
    "gopkg.in/yaml.v2"
)

type Server struct {
    PHP		string	`yaml:"php"`
    JAVA    string  `yaml:"java"`
    KOTLIN	string  `yaml:"kotlin"`
    NODE	string  `yaml:"node"`
    RUBY	string  `yaml:"ruby"`
    PYTHON	string	`yaml:"python"`
    GO		string	`yaml:"go"`
}

type Config struct {
    Server *Server  `yaml:"server"`
    Port    int     `yaml:"port"`
}

func New(filename string) (*Config, error) {
    var config *Config
    
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        return config, err
    }
    
    err = yaml.Unmarshal(bytes, &config)
    if err != nil {
        return config, err
    }
    
    return config, nil
}