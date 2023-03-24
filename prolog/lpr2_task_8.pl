student("Миша", "5а").
student("Маша", "6б").
student("Леша", "6в").
student("Рома", "7в").
student("Таня", "8а").


hobby("Миша", "кино").
hobby("Маша", "кино").
hobby("Леша", "кино").
hobby("Рома", "лыжи").
hobby("Таня", "лыжи").


show_all :-
    student(N, C),
    hobby(N, H),
    write(N),
    write(" "),
    write(C),
    write(" "),
    write(H).

cinema(C1) :-
    student(N, C1),
    hobby(M, "кино"),
    student(M, C2),
    C1\==C2,
    write(N),
    write(" "),
    write(M),
    write(" "),
    write(C2).
