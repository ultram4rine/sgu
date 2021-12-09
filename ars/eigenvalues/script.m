N=100000;
m=10;

Diags=zeros(N,5);
Diags(1:N,1)=10*sqrt(1+(1:N));
Diags(1:N-1,2)=1i;
Diags(1:N-1,3)=1;
Diags(1:N-2,4)=1/4;
Diags(1:N-2,5)=-1/4;

tic;
A=spdiags(Diags, [0,1,-1,2,-2], N, N);
A(N,1)=-1;
A(1,N)=-1i;
%disp(A);
toc;

tic;
Lam=eigs(A, m, 'sm');
toc;

disp('Собственные значения');
disp(Lam);
clear all;
