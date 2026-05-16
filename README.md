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

</div>

<br>

## 💾 [ SOBRE O SISTEMA ]

Este é um projeto desenvolvido inteiramente na linguagem **Go (Golang)**, focado em aplicar os conceitos de **Orientação a Objetos (OOP)** em um ambiente que, nativamente, não possui classes. Através deste sistema em CLI (Command Line Interface), é possível gerenciar contas bancárias, clientes, realizar pagamentos de boletos, saques, depósitos e transferências.

================================================================================

## ⚙️ [ HARDWARE / TECNOLOGIAS UTILIZADAS ]

Neste programa de simulação bancária, foram empregadas as seguintes diretrizes e ferramentas da linguagem Go:

*   **`Structs` (Estruturas de Dados):** Usadas como os pilares do sistema (substituindo o conceito clássico de "Classes"). 
    *   `Titular`: Armazena dados do cliente (Nome, CPF, Profissão).
    *   `ContaCorrente` / `ContaPoupanca`: Herdam (através de Composição) os dados do Titular e mantêm o controle de agência, conta e `saldo` (encapsulado e protegido para não ser modificado diretamente).

*   **`Métodos` com Ponteiros (`*`):** As funções de `Depositar`, `Sacar` e `Transferir` utilizam ponteiros de receiver (ex: `c *ContaCorrente`). Isso permite que o método altere o saldo real alocado na memória, em vez de alterar apenas uma cópia descartável da conta.

*   **`Interfaces` (Polimorfismo):** A mágica do sistema! Foi criada a interface `verificarConta`, que estabelece um contrato: _"Qualquer estrutura que tenha um método `Sacar(valor, operacao)` pode ser considerada uma conta válida"_. Graças a isso, a função `PagarBoleto` consegue aceitar tanto `ContaCorrente` quanto `ContaPoupanca` sem precisar duplicar código.

*   **Módulo Padrão `fmt`:** Para toda a interface de saída (Input/Output). Foi fortemente utilizado o `fmt.Println` para feedbacks no terminal e `fmt.Sprintf` para formatar strings de extrato bancário saltando linhas (`\n`) de forma limpa e organizada.

================================================================================

## 🕹️ [ MANUAL DE OPERAÇÃO (COMO RODAR) ]

**PRÉ-REQUISITOS:**
* O terminal precisa ter o compilador do Go instalado. (Testado na versão 1.26.2).

**BOOT DO SISTEMA:**
1. Clone o repositório para o seu disco rígido:
   ```bash
   git clone https://github.com/SEU_USUARIO/SEU_REPOSITORIO.git
   ```
2. Acesse o diretório do banco pelo console:
   ```bash
   cd bank
   ```
3. Execute a rotina principal:
   ```bash
   go run main.go
   ```

================================================================================

## 🖨️ [ OUTPUT ESPERADO ]

Ao executar o programa, o sistema fará depósitos e debitará boletos, exibindo um extrato semelhante a este monitor CRT:

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

================================================================================

<div align="center">
  <i>Desenvolvido durante os estudos da Trilha Go. End of File (EOF).</i> 👾
</div>
