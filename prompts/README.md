# Prompts Folder

Store in this folder the prompts that were used during code generation, analysis, refactoring, and test creation for this kata.

Each prompt file should document the actual prompt text that was used, so the work can be reviewed and the prompting process can be reproduced.

1. we have to work on this repository to change the current code maintaining it's behaviour, but symplifying it. We will work only on the rust code, keepng go and java code out of our scope. The first thing i want to ask you is to summarize what is in the lib file, to understand exaclty what have already been done and why it is hard to undertand, maintain and extend

2. Ok, i have understood the situation. What we have to do now is to refactor all the code in order to simplify it.
The important thing is to MAINTAIN the behaviour. Everything must work as it does now, but with different implementation.
Let's change the implementation to make it extendable and maintainable. Rewrite the if/else statement, adjust the stringly-typed domain and remove the magic numbers. Add separate function for better code understanding and management.