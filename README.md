Supports instructions:

| Instruction | Args |
| --- | --- |
| ADD | {Rd,} Rn, Op  |
| SUB | {Rd,} Rn, Op |
| MUL | {Rd,} Rn, Rm |
| B | label |
| MOV | Rd, Op |
| SDIV | {Rd,} Rn, Rm |

All instructions support conditional suffixes, such as 
| Suffix | Condition |
| --- | --- |
| EQ | Z = 1 |
| NE | Z = 0 |
| CS/HS | C = 1 |
| CC/LO | C = 0 |
| MI | N = 1 |
| PL | N = 0 |
| VS | V = 1 |
| VC | V = 0 |
| HI | C = 1 && Z = 0 |
| LS | C = 0 && Z = 1 |
| GE | N = V |
| LT | N != V | 
| GT | Z = 0 && N = V |
| LE | Z = 1 && N != V |
| AL | Always |

Instructions **ADD**, **SUB** **AND** **MUL** support suffix **S**, which sets state flags **Z**, **N**, **V**, **C** according to result.

Supoorts labels described as `label_name:`
