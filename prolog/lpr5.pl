lr("Миша", "Юра").
lr("Витя", "Миша").

ord(F, S, T) :-
    lr(F, S),
    lr(S, T).

higher("ель", "сосна").
higher("сосна", "тополь").
higher("тополь", "береза").
higher("береза", "липа").
higher("липа", "клен").

hl(A, B, C, D, E, F) :-
    higher(A, B),
    higher(B, C),
    higher(C, D),
    higher(D, E),
    higher(E, F),
    write("higher "),
    write(A),
    nl,
    write("lower "),
    write(F).

name("Петя").
name("Лена").
name("Таня").

pet("собака").
pet("кошка").
pet("хомячек").

man("Петя").
woman("Лена").
woman("Таня").

map(X, Y) :-
    name(X),
    pet(Y),
    X="Таня",
    Y="кошка".
map(X, Y) :-
    name(X),
    pet(Y),
    Y="хомячек",
    woman(X),
    not(X="Таня").
map(X, Y) :-
    name(X),
    pet(Y),
    X="Петя",
    not(Y="кошка").

solution(X1, Y1, X2, Y2, X3, Y3) :-
    X1="Таня",
    map(X1, Y1),
    X2="Лена",
    map(X2, Y2),
    X3="Петя",
    map(X3, Y3),
    not(Y1=Y2),
    not(Y2=Y3),
    not(Y1=Y3).
