likes("Миша", "гитара").
likes("Маша", "арфа").
likes("Рома", "футбол").
likes("Таня", "лыжи").

sport("футбол").
sport("лыжи").

musical_instrument("арфа").
musical_instrument("гитара"). 

musician(Кто) :-
    likes(Кто, Чем),
    musical_instrument(Чем).

vacation("Саша", "Анталия").
vacation("Анна", "Сочи").
vacation("Миша", "Юрмала").
vacation("Коля", "Рим").

italy("Рим").

russia("Сочи").

baltic("Юрмала").

vacation_russia(Кто) :-
    vacation(Кто, Где),
    russia(Где).
