# Dabulang (ダブ言語)

Dabulang is an interpreted object-oriented programming language aimed towards game development. The language's standard library has a wrapper around [raylib](https://www.raylib.com) so, developing games in dabulang is a fairly simple process.

## Design

Dabulang has a similar syntax to BASIC, drawing inspiration from the TI's version of BASIC: [TI-BASIC](https://en.wikipedia.org/wiki/TI-BASIC). In a nutshell, blocks of statements/expressions need to end with the `End` keyword, and all keywords start with a capital letter.

In addition to TI-BASIC, Dabulang has classes, functions, singletons and a syntax that is more suited for programming on a computer as opposed to the original TI graphic calculator.

## Impression

```Dabulang
Fold Person
  name, age
End

Func greet(p)
  print("Hello, ");
  println(p(name));
End

john = Person("John", 35);
mary = Person("Mary", 33);

people = [john, mary];

For (i = 0; i < len(people); i = i + 1)
  greet(people(i));
End
```
