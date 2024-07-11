# Estruturas de Dados e Algoritmos (EDA)

*[Read in English](README.md)*

Este projeto tem como objetivo demonstrar as principais estruturas de dados e algoritmos, bem como documentar e explicar cada um deles.

## Estruturas de Dados

### Array (Vetor)

Um Array (Vetor) é uma estrutura de dados que armazena uma coleção de elementos em locais de memória contíguos. Cada elemento pode ser acessado diretamente pelo seu índice, tornando as operações de acesso muito rápidas. No entanto, a inserção e remoção de elementos podem ser demoradas, pois podem exigir o deslocamento de outros elementos.

**Características:**
- **Indexação**: Acesso direto aos elementos via índice.
- **Tamanho fixo**: O tamanho do vetor é definido na criação e não pode ser alterado.
- **Tipo homogêneo**: Todos os elementos do vetor são do mesmo tipo.

**Principais operações e suas complexidades:**
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
A implementação da estrutura de dados Array em Go pode ser encontrada no arquivo [`array.go`](src/data_structures/linear/array.go).

### Lista Ligada (Linked List)

Uma Lista Ligada é uma estrutura de dados linear onde cada elemento é um nó que contém um valor e uma referência (ou ponteiro) para o próximo nó na sequência.

**Características:**
- **Cadeia de Nós**: Cada nó aponta para o próximo nó.
- **Tamanho Variável**: A lista pode crescer e encolher dinamicamente.
- **Acesso Sequencial**: O acesso aos elementos é sequencial, começando pelo primeiro nó.

**Principais Operações e Suas Complexidades:**
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
A implementação da estrutura de dados Lista Ligada em Go pode ser encontrada no arquivo [`linked_list.go`](src/data_structures/linear/linked_list.go).

### Pilha (Stack)

Uma Pilha é uma estrutura de dados linear que segue o princípio de Último a Entrar, Primeiro a Sair (LIFO). Os elementos são adicionados e removidos da mesma extremidade, chamada de topo da pilha.

**Características:**
- **LIFO**: Ordenação de Último a Entrar, Primeiro a Sair.
- **Tamanho Variável**: A pilha pode crescer e encolher dinamicamente.
- **Acesso ao Topo**: O acesso aos elementos é somente no topo da pilha.

**Principais Operações e Suas Complexidades:**
- **Push** (adicionar um elemento ao topo):
  - O(1) - tempo constante.
- **Pop** (remover um elemento do topo):
  - O(1) - tempo constante.
- **Peek** (visualizar o elemento do topo sem removê-lo):
  - O(1) - tempo constante.
- **Tamanho**:
  - O(1) - tempo constante.
- **EstáVazia**:
  - O(1) - tempo constante.

**Implementação em Go:**
A implementação da estrutura de dados Pilha em Go pode ser encontrada no arquivo [`stack.go`](src/data_structures/linear/stack.go).

### Fila (Queue)

Uma Fila é uma estrutura de dados linear que segue o princípio de Primeiro a Entrar, Primeiro a Sair (FIFO). Os elementos são adicionados na parte traseira (traseira) e removidos da frente.

**Características:**
- **FIFO**: Ordenação de Primeiro a Entrar, Primeiro a Sair.
- **Tamanho Variável**: A fila pode crescer e encolher dinamicamente.
- **Acesso Frontal e Traseiro**: Os elementos são adicionados na parte traseira e removidos da parte frontal.

**Principais Operações e Suas Complexidades:**
- **Enqueue** (adicionar um elemento à parte traseira):
  - O(1) - tempo constante.
- **Dequeue** (remover um elemento da parte frontal):
  - O(1) - tempo constante.
- **Peek** (visualizar o elemento da frente sem removê-lo):
  - O(1) - tempo constante.
- **Tamanho**:
  - O(1) - tempo constante.
- **EstáVazia**:
  - O(1) - tempo constante.

**Implementação em Go:**
A implementação da estrutura de dados Fila em Go pode ser encontrada no arquivo [`queue.go`](src/data_structures/linear/queue.go).

### Deque (Double-ended Queue)

Um Deque (Fila de Dupla Extremidade) é uma estrutura de dados linear que permite a inserção e remoção de elementos de ambas as extremidades, frente e traseira.

**Características:**
- **Dupla Extremidade**: Os elementos podem ser adicionados ou removidos de ambas as extremidades.
- **Tamanho Variável**: O deque pode crescer e encolher dinamicamente.
- **Acesso Frontal e Traseiro**: Acesso aos elementos é na frente e na traseira do deque.

**Principais Operações e Suas Complexidades:**
- **AdicionarFrente**:
  - O(n) - tempo linear.
- **AdicionarTraseira**:
  - O(1) - tempo constante.
- **RemoverFrente**:
  - O(1) - tempo constante.
- **RemoverTraseira**:
  - O(1) - tempo constante.
- **Tamanho**:
  - O(1) - tempo constante.
- **EstáVazio**:
  - O(1) - tempo constante.

**Implementação em Go:**
A implementação da estrutura de dados Deque em Go pode ser encontrada no arquivo [`deque.go`](src/data_structures/linear/deque.go).

### Árvore (Tree)

Uma Árvore é uma estrutura de dados hierárquica composta por nós, onde cada nó possui um valor e referências para seus filhos esquerdo e direito.

**Características:**
- **Estrutura Hierárquica**: Composta por nós com relacionamentos pai-filho.
- **Nó Raiz**: O nó superior com zero ou mais nós filhos.
- **Subárvores**: Cada nó filho pode ter seus próprios filhos, formando subárvores.
- **Árvores Binárias**: Árvores onde cada nó tem no máximo dois filhos (esquerdo e direito).

**Principais operações e suas complexidades:**
- **Inserir**:
  - Caso médio: O(log n) - tempo logarítmico.
  - Pior caso: O(n) - tempo linear (árvore desbalanceada).
- **Buscar**:
  - Caso médio: O(log n) - tempo logarítmico.
  - Pior caso: O(n) - tempo linear (árvore desbalanceada).
- **Remover**:
  - Caso médio: O(log n) - tempo logarítmico.
  - Pior caso: O(n) - tempo linear (árvore desbalanceada).
- **EstáVazia**:
  - O(1) - tempo constante.

**Implementação em Go:**
A implementação da estrutura de dados Árvore em Go pode ser encontrada no arquivo [`tree.go`](src/data_structures/tree-based/tree.go).

## Algoritmos

(Em breve)
