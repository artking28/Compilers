# Mantis

**Mantis** é uma linguagem de programação simples criada como trabalho de compiladores do SENAC, ela contém um CLI
enxuto
para compilar, rodar e validar arquivos
mantis(.
mnts) para bytecode neander.

## Comandos

```bash
  mantis build arquivo.mnts;
```
Compila o arquivo Mantis.

```bash
  mantis run arquivo.mnts
```
Roda o arquivo Mantis.

```bash
  mantis vet arquivo.mnts
```
Verifica a validade do arquivo Mantis.

## Flags

- `-s <subset>`, `--subset=<subset>` **[opcional]** - Define o subset a ser usado para compilar o código.
- `-o <file>` **[opcional]** - Define o nome e tipo do output gerado.
