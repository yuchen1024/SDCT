## Specifications

- OS: MAC OS x64

- Language: C++

- Requires: OpenSSL

- The default elliptic curve is "NID_X9_62_prime256v1"

---

## Installation

The current implementation is based on OpenSSL library. See the installment instructions of OpenSSL as below:  

1. Download [openssl-master.zip](https://github.com/openssl/openssl.git)

2. make a directory "openssl" to save the source codes

```
    mkdir openssl
    mv openssl-master.zip /openssl
```

3. unzip it

4. install openssl on your machine

```
    ./config --prefix=/usr/local/ssl shared
    make 
    sudo make install
```

---

## Code Structure of SDCT_Cryptocurrency

/src: SDCT.cpp: algorithms of SDCT 

/depends: depends module

- /common: common module 
  * global.hpp: generate global parameters
  * print.hpp: print split line for demo use
  * hash.hpp: case-tailored hash functions based on SHA2 and SM3
  * routines.hpp: related routine algorithms

- /twisted_elgamal: PKE module
  * twisted_elgamal.hpp: twisted ElGamal PKE  
  * calculate_dlog.hpp: discrete logarithm searching algorithm

- /nizk: Sigma-protocols  
  * nizk_plaintext_equality.hpp: NIZKPoK for triple twisted ElGamal plaintext equality
  * nizk_plaintext_knowledge.hpp: NIZKPoK for twisted ElGamal plaintext and randomness knowledge
  * nizk_dlog_equality.hpp: NIZKPoK for discrete logarithm equality

- /bulletproofs: Bulletproofs module
  * aggregate_bulletproof.hpp: the aggregating logarithmic size Bulletproofs
  * innerproduct_proof.hpp: the inner product argument (used by Bulletproof to shrink the proof size) 

/test: all test files

---

## Compile and Run

To compile and run the SDCT system, do the following: 

```
  $ mkdir build && cd build
  $ cmake ..
  $ make -j8
  $ ./test_SDCT
```
---

## Test

set the range size = $[0, 2^\ell = 2^{32}-1 = 4294967295]$

### Flow of SDCT_Cryptocurrency

   1. run <font color=blue>Setup</font> to build up the system, 
      generating system-wide parameters, store public parameters into "SDCT.pp", 
      store secret parameters into "SDCT.sp"
   2. run <font color=blue>CreateAccount</font> to create accounts for Alice ($m_1$) and Bob ($m_2$); 
      one can reveal the balance by running <font color=blue>Reveal_Balance:</font> 
   3. Alice runs <font color=blue>CreateCTx</font> to transfer $v_1$ coins to Bob ===> Alice_sn.ctx; 
      <font color=blue>Print_CTx:</font> shows the details of CTx
   4. Miners runs <font color=blue>VerifyCTx</font> check CTx validity
   5. If CTx is valid, run <font color=blue>UpdateAccount</font> 
   to update Alice and Bob's account balance and serialize the changes.


### Test Cases
---
Create SDCT environment

1. setup the SDCT system, generates $pp$ and $sp$

2. generate two accounts: Alice and Bob
   * $512$ --- Alice's initial balance  
   * $256$ --- Bob's initial balance    

3. Invalid CTx: <font color=red>$v_1 \neq v_2$ $\Rightarrow$ plaintext equality proof will be rejected</font>  
   - $v_1 \neq v_2$ in transfer amount

4. Invalid CTx: <font color=red>$v \notin [0, 2^\ell]$ $\Rightarrow$ range proof for right interval will be rejected</font>
   - $v  = 4294967296 > 4294967295$  

5. Invalid CTx: <font color=red>$(m_1 - v) \notin [0, 2^\ell]$ $\Rightarrow$ range proof for solvent 
   will be rejected</font>
   - $m_1  = 512$ --- Alice's balance  
   - $v  = 513$ --- transfer amount 

6. 1st Valid CTx: <font color=blue>$v_1 = v_2 \wedge v_1 \in [0, \ell] \wedge (m_1 - v_1) \in [0, \ell]$</font>
   - $v    = 128$ --- transfer amount from Alice to Bob
   - $384$ --- Alice's updated balance  
   - $384$ --- Bob's updated balance    

7. 2nd Valid CTx: <font color=blue>$v_1 = v_2 \wedge v_1 \in [0, \ell] \wedge (m_1 - v_1) \in [0, \ell]$</font>
   - $v    = 384$ --- transfer amount from Bob to Alice
   - $768$ --- Alice's updated balance  
   - $0$ --- Bob's updated balance    

8. Supervisor opens ctx1
   - $v = 128$  

9. Supervisor opens ctx2: 
   - $v = 384$   

---