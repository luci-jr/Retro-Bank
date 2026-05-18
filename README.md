<div align="center">

```text
  ____  _____ _____ ____   ___    ____    _    _   _ _  __
 |  _ \| ____|_   _|  _ \ / _ \  | __ )  / \  | \ | | |/ /
 | |_) |  _|   | | | |_) | | | | |  _ \ / _ \ |  \| | ' / 
 |  _ <| |___  | | |  _ <| |_| | | |_) / ___ \| |\  | . \ 
 |_| \_\_____| |_| |_| \_\\___/  |____/_/   \_\_| \_|_|\_\

 :: SISTEMA BANCÁRIO v1.0 - TERMINAL EDITION ::
```

> _"Acesse o mainframe bancário e gerencie suas finanças em Go!"_ 📟

---

```text
🎮 [ B O O T _ S E Q U E N C E ]
> INICIANDO_SISTEMA.................... [OK]
> CARREGANDO_MODULOS_GO............... [OK]
> VERIFICANDO_INTEGRIDADE_OOP......... [OK]
> SISTEMA_PRONTO.
```

</div>

<br>

## 💾 [ SOBRE O SISTEMA ]

Este é um projeto desenvolvido inteiramente na linguagem **Go (Golang)**, focado em aplicar os conceitos de **Orientação a Objetos (OOP)** em um ambiente que, nativamente, não possui classes. Através deste sistema em CLI (Command Line Interface), você pode simular as operações de um banco real diretamente no terminal.

---

## 🕹️ [ FUNCIONALIDADES ]

- [x] **Criação de Contas**: Suporte a Conta Corrente e Conta Poupança.
- [x] **Operações Básicas**: Saques, depósitos e consultas de saldo.
- [x] **Transferências**: Envio de valores entre contas.
- [x] **Pagamento de Boletos**: Utilização de Interfaces para permitir pagamentos a partir de qualquer tipo de conta válida.
- [x] **Encapsulamento**: Proteção do saldo contra manipulações diretas.

---

## ⚙️ [ RECURSOS DO GO UTILIZADOS ]

Neste programa de simulação bancária, foram empregadas as seguintes diretrizes e ferramentas da linguagem Go:

*   **`Structs` (Estruturas de Dados):** Usadas como os pilares do sistema (substituindo o conceito clássico de "Classes").
    *   `Titular`: Armazena dados do cliente (Nome, CPF, Profissão).
    *   `ContaCorrente` / `ContaPoupanca`: Herdam (através de Composição) os dados do Titular e mantêm o controle de agência, conta e `saldo`.
*   **`Métodos` com Ponteiros (`*`):** As funções de `Depositar`, `Sacar` e `Transferir` utilizam ponteiros de receiver. Isso permite que o método altere o saldo real alocado na memória.
*   **`Interfaces` (Polimorfismo):** Foi criada a interface `verificarConta`, permitindo que a função `PagarBoleto` aceite tanto `ContaCorrente` quanto `ContaPoupanca` sem duplicar código.
*   **Módulo Padrão `fmt`:** Utilizado para toda a interface de I/O, formatação de strings e exibição de extratos.

---

## 🚀 [ MANUAL DE OPERAÇÃO ]

### Pré-requisitos
* Ter o compilador do **Go** instalado em sua máquina.

### Como Executar

**1. Clone o repositório:**
```bash
git clone https://github.com/lucivaldo/Retro-Bank.git
```

**2. Acesse o diretório do projeto:**
```bash
cd Retro-Bank
```

**3. Execute a aplicação:**
```bash
go run main.go
```

---

## 🖨️ [ OUTPUT ESPERADO ]

Ao executar o programa, o sistema exibirá as operações no terminal como em um monitor CRT antigo:

```text
Deposito realizado com sucesso!
Saldo atual: R$ 1000
Pagamento de boleto realizado com sucesso!
Saldo atual: R$ 500
--- Saldo final: R$ 500 ---

Deposito realizado com sucesso!
Saldo atual: R$ 2000
Saldo insuficiente para Pagamento de boleto 
O saldo se mantem em: R$ 10
--- Consulta final: R$ 10 ---
```

---

<div align="center">
  <i>Desenvolvido durante os estudos da Trilha Go. End of File (EOF).</i> 👾
</div>
