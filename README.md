# go-boilerplate

Esse aplicativo é um modelo de como organizar aplicações usando Golang.
A ideia é ter uma aplicação em camadas isoladas, pouco acoplada e testável.

Não é uma aplicação finalizada, mas sim uma forma de praticar boas práticas de programação.

As funcionalidades do _software_ é registrar e controlar uma biblioteca pessoal.

## Arquitetura

A aplicação tem três principais camadas: **handlers**, **services** e **repositórios**. A comunicação entre camadas se da por interfaces que são dependências injetadas.

A ideia é que assim, as camadas possam ser desenvolvidas isoladamente sem quebrar a aplicação. Além disso, facilita os testes unitários.