# Dabulang (ダブ言語)

Dabulang is an interpreted object-oriented programming language aimed towards game development. The language's standard library has a wrapper around [raylib](https://www.raylib.com) so, developing games in dabulang is a fairly simple process.

## Design

Dabulang has a similar syntax to BASIC, drawing inspiration from the TI's version of BASIC: [TI-BASIC](https://en.wikipedia.org/wiki/TI-BASIC). In a nutshell, blocks of statements/expressions need to end with the `End` keyword, and all keywords start with a capital letter.

In addition to TI-BASIC, Dabulang has classes, functions, singletons and a syntax that is more suited for programming on a computer as opposed to the original TI graphic calculator.

## Impression

```Dabulang
Singleton Player Extends Sprite
  Prop name = "", age;

  Func init()
    Super(10, 10, "player.png");
  End

  Func hello(a, b)
    If a == b Then
      matrix = [[[1, 2], [3, 4], [5, 6]], [[7, 8], [9, 10], [11, 12]]];
      element = G(0, 2, 0);
      Return element;
    Else
      Return 10;
    End
  End
End

Player.hello(20, Player.name);
```
