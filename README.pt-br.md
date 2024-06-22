# Estruturas de Dados e Algoritmos (DSA)

*[Read in English](README.md)*

Este projeto tem como objetivo demonstrar as principais estruturas de dados e algoritmos, bem como documentar e explicar cada um deles.

## Data Structures

### Array

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

### Lista Ligada/Encadeada

Uma Lista Ligada/Encadeada (Linked List) é uma estrutura de dados linear onde cada elemento é um nó que contém um valor e uma referência (ou ponteiro) para o próximo nó na sequência.

**Características:**
- **Cadeia de Nós**: Cada nó aponta para o próximo nó.
- **Tamanho Variável**: A lista pode crescer e diminuir dinamicamente.
- **Acesso Sequencial**: O acesso aos elementos é sequencial, começando pelo primeiro nó.

**Operações principais e suas complexidades:**
- **Acesso**:
  - Melhor caso: O(1) - tempo constante (primeiro elemento).
  - Pior caso: O(n) - tempo linear (último elemento).
- **Inserção**:
  - Melhor caso: O(1) - tempo constante (inserção no início).
  - Pior caso: O(n) - tempo linear (inserção no final).
- **Remoção**:
  - Melhor caso: O(1) - tempo constante (remoção do início).
  - Pior caso: O(n) - tempo linear (remoção do final).

**Implementação em Go:**
A implementação da estrutura de dados Linked List em Go pode ser encontrada no arquivo [`linked_list.go`](src/data_structures/linked_list.go).

### Pilha

Uma Pilha (Stack) é uma estrutura de dados linear que segue o princípio de Último a Entrar, Primeiro a Sair (LIFO). Elementos podem ser adicionados e removidos apenas do topo da pilha.

**Características:**
- **LIFO**: Ordenação de Último a Entrar, Primeiro a Sair.
- **Tamanho Variável**: A pilha pode crescer e diminuir dinamicamente.
- **Acesso ao Topo**: O acesso aos elementos é feito apenas no topo da pilha.

**Operações principais e suas complexidades:**
- **Push**:
  - O(1) - tempo constante.
- **Pop**:
  - O(1) - tempo constante.
- **Peek**:
  - O(1) - tempo constante.
- **Size**:
  - O(1) - tempo constante.
- **IsEmpty**:
  - O(1) - tempo constante.

**Implementação em Go:**
A implementação da estrutura de dados Pilha em Go pode ser encontrada no arquivo [`stack.go`](src/data_structures/stack.go).

### Fila

Uma Fila (Queue) é uma estrutura de dados linear que segue o princípio de Primeiro a Entrar, Primeiro a Sair (FIFO). Elementos podem ser adicionados no final e removidos do início da fila.

**Características:**
- **FIFO**: Ordenação de Primeiro a Entrar, Primeiro a Sair.
- **Tamanho Variável**: A fila pode crescer e diminuir dinamicamente.
- **Acesso ao Início e ao Fim**: O acesso aos elementos é feito apenas no início e no fim da fila.

**Operações principais e suas complexidades:**
- **Enqueue**:
  - O(1) - tempo constante.
- **Dequeue**:
  - O(1) - tempo constante.
- **Peek**:
  - O(1) - tempo constante.
- **Size**:
  - O(1) - tempo constante.
- **IsEmpty**:
  - O(1) - tempo constante.

**Implementação em Go:**
A implementação da estrutura de dados Fila em Go pode ser encontrada no arquivo [`queue.go`](src/data_structures/queue.go).

## Algorithms

(Em breve)
