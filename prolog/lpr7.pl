:- (dynamic true/2, false/2).

target :-
    lang(A), !,
    write("Загаданный язык программирования "),
    writeln(A),
    clear.
target :-
    write("К сожалению, этот ЯП мне неизвестнен."),
    clear.

lang("C") :-
    is("компилируемый"),
    is("императивный"),
    attr("На нем", "написано ядро Linux").
lang("C++") :-
    is("компилируемый"),
    is("объектно-ориентированный"),
    attr("Его изобрел", "Страуструп").
lang("Golang") :-
    is("компилируемый"),
    is("многопоточный"),
    is("императивный"),
    attr("У него есть", "горутины").
lang("Python") :-
    is("интерпретируемый"),
    is("объектно-ориентированный"),
    attr("Он", "назван в честь змеи").
lang("PHP") :-
    is("интерпретируемый"),
    is("объектно-ориентированный"),
    attr("На нем", "сделан Facebook").
lang("Javascript") :-
    is("интерпретируемый"),
    is("объектно-ориентированный"),
    attr("Он", "работает в браузере").
lang("Java") :-
    is("компилируемый"),
    is("многопоточный"),
    is("объектно-ориентированный"),
    attr("Он", "имеет виртуальную машину").
lang("Rust") :-
    is("компилируемый"),
    is("императивный"),
    is("функциональный"),
    attr("У него есть", "механизм владения").
lang("Prolog") :-
    is("интерпретируемый"),
    is("логический"),
    attr("Он", "основан на предикатах Хорна").

is("компилируемый") :-
    attr("У него есть", "компилятор").
is("интерпретируемый") :-
    attr("У него есть", "интерпретатор").
is("многопоточный") :-
    attr("Он", "умеет параллельно обрабатывать данные").
is("императивный") :-
    attr("Он", "последовательно выполняет инструкции").
is("объектно-ориентированный") :-
    attr("У него есть", "объекты").
is("функциональный") :-
    attr("Он", "основан на чистоте функций").
is("логический") :-
    attr("Он", "задается логическими утверждениями").

attr(A, B) :-
    true(A, B), !.
attr(A, B) :-
    not(false(A, B)),
    know(A, B, Ans),
    Ans=y.

no_attr(A, B) :-
    false(A, B), !.
no_attr(A, B) :-
    not(true(A, B)),
    know(A, B, Ans),
    Ans=n.

know(A, B, Ans) :-
    write(A),
    write(" "),
    write(B),
    read(Ans),
    remember(A, B, Ans).

remember(A, B, y) :-
    assert(true(A, B)).
remember(A, B, n) :-
    assert(false(A, B)).

clear :-
    retract(true(_, _)),
    fail.
clear :-
    retract(false(_, _)),
    fail.
