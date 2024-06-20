# Estruturas de Dados e Algoritmos (DSA)

*[Read in English](README.md)*

Este projeto tem como objetivo demonstrar as principais estruturas de dados e algoritmos, bem como documentar e explicar cada um deles.

## Data Structures

### Array

[W3S - Array](https://www.w3schools.com/dsa/dsa_data_arrays.php)

Um Array é uma estrutura de dados que armazena uma coleção de elementos em locais de memória contíguos. Cada elemento pode ser acessado diretamente pelo seu índice, o que torna a operação de acesso muito rápida. No entanto, a inserção e a remoção de elementos podem ser demoradas, pois podem exigir o deslocamento de outros elementos.

**Características:**
- **Indexação**: Acesso direto aos elementos via índice.
- **Tamanho fixo**: O tamanho do array é definido na sua criação e não pode ser alterado.
- **Tipo homogêneo**: Todos os elementos do array são do mesmo tipo.

**Operações principais e suas complexidades:**
- **Acesso**:
  - Melhor caso: O(1) - tempo constante.
  - Pior caso: O(1) - tempo constante.
- **Inserção**:
  - Melhor caso: O(1) - tempo constante (inserção no final se houver espaço).
  - Pior caso: O(n) - tempo linear (realocação de elementos).
- **Remoção**:
  - Melhor caso: O(1) - tempo constante (remoção do último elemento).
  - Pior caso: O(n) - tempo linear (deslocamento de elementos).

**Implementação em Go:**
A implementação da estrutura de dados Array em Go pode ser encontrada no arquivo [`array.go`](src/data_structures/array.go).

## Algorithms

(Em breve)
