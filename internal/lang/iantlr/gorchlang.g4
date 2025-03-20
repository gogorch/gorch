grammar gorchlang;

options{
    language=Go;
}

// 入口描述
primary
    : (startDirective | fragmentDirective | registerDirective | registerDirective)*
;

// START子描述，执行入口
startDirective
    : 'START' '(' (name = stringLiteral) (',' argsStmt)? ')'
    '{'
    ((onFinishStmt noCheckMissDirective) | (noCheckMissDirective onFinishStmt) | noCheckMissDirective | onFinishStmt)?
    exedescStmt
    '}'
;

// FRAGMENT子描述，可以被START描述引用
fragmentDirective
    : 'FRAGMENT' '(' (name = stringLiteral) ')' '{'
    exedescStmt
    '}'
;

// 注册子描述，用来注册算子和选项
registerDirective
    : 'REGISTER' '(' (pkg = stringLiteral) ')' '{'
    registerOperatorDirective *
    '}'
;

// START包含的结束事件
onFinishStmt
    : 'ON_FINISH' '(' ')' '{'
    exedescStmt
    '}'
;

exedescStmt
    : serialStmt
    | concurrentStmt
    | wrapStmt
    | wrapBracketStmt
    | operatorStmt
    | switchDirective
    | unfoldDirective
;

// 执行片段
leafSnippet
    : operatorStmt
    | concurrentStmt
    | serialBracketStmt
    | unfoldDirective
    | wrapBracketStmt
    | goDirective
    | switchDirective
;

operatorStmt
    : ignoreError? (operatorName = ID)  // 啥也不带
    | ignoreError? (operatorName = ID) '(' ')' // 只带一个括号
    | ignoreError? (operatorName = ID) ('(' argsStmt ')')? // 只带参数
    | ignoreError? (operatorName = ID) ('(' waitDirective (',' waitDirective)* ')')? // 只带事件
    | ignoreError? (operatorName = ID) ('(' argsStmt ',' waitDirective (',' waitDirective)* ')')? // 事件和参数都带
;

serialStmt
    : (leafSnippet | skipDirective) ('->' (leafSnippet | skipDirective))+
;

serialBracketStmt
    : '(' serialStmt ')'
;

concurrentStmt
    : '[' leafSnippet (',' leafSnippet)* ','? ']'
;

unfoldDirective
    : 'UNFOLD' '(' (name = stringLiteral) ')'
;

goDirective
    : 'GO' '(' exedescStmt ',' (name = stringLiteral) ')'
;

waitDirective
    : ignoreError? 'WAIT' '(' (name = stringLiteral) (',' argsStmt)? ')'
;

// START指令内，不校验算子是否缺失
noCheckMissDirective
    : 'NO_CHECK_MISS' '(' ')'
;

// 串行组支持跳过
skipDirective
    : 'SKIP' '(' operatorStmt ')'
;

// 包装算子语法
wrapStmt
    : operatorStmt ('|' operatorStmt)* '|' leafSnippet
;

// 作为叶子节点的包装算子，需要加上括号
wrapBracketStmt
    : '(' wrapStmt ')'
;

switchDirective
    : 'SWITCH' '(' operatorStmt ')' '{' switchCaseDirective (',' switchCaseDirective)* ','? '}'
;

switchCaseDirective
    : 'CASE' (caseName = stringLiteral) '=>' leafSnippet
;

registerOperatorDirective
    : 'OPERATOR' '(' 
    (((pkgPath = stringLiteral) ',' (structName = stringLiteral) ',' (operatorName = stringLiteral) ',' (seq = integerLiteral)) |
    ((pkgPath = stringLiteral) ',' (structName = stringLiteral) ',' (seq = integerLiteral)))
    ')'
;

argStmt
    : (key = ID) '=' boolLiteral
    | (key = ID) '=' integerLiteral
    | (key = ID) '=' stringLiteral
    | (key = ID) '=' durationLiteral
    | (key = ID) '=' '[' ']'
    | (key = ID) '=' '[' boolLiteral (',' boolLiteral)* ']'
    | (key = ID) '=' '[' integerLiteral (',' integerLiteral)* ']'
    | (key = ID) '=' '[' stringLiteral (',' stringLiteral)* ']'
    | (key = ID) '=' '[' durationLiteral (',' durationLiteral)* ']'
;

argsStmt
    : argStmt (',' argStmt)*
;

ignoreError
    : '@'
;

integerLiteral: '-'? INT;
stringLiteral: DQUOTA_STRING;
boolLiteral: TRUE | FALSE;
durationLiteral: INT ('s'|'ms'|'µs'|'h'|'m');

TRUE: 'TRUE' | 'true' | 'True';
FALSE: 'FALSE' | 'false' | 'False';

ID:  ('a'..'z' |'A'..'Z'| '_')+ ( ('0'..'9') | ('a'..'z' |'A'..'Z') | '_' )* ;

INT: '0'..'9' +;

DQUOTA_STRING: '"' ( '\\'. | '""' | ~('"' | '\\') )* '"';
LINE_COMMENT: '//' ~[\r\n]* -> skip;
COMMENT: '/*' .*? '*/' -> skip;
TERMINATOR: [\r\n]+  -> skip;
WS:[ \t\n\r]+ -> skip;