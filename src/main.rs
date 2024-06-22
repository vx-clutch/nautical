use clap::{Parser, ValueEnum};

enum Token {
    IDENT,
    IntLit,
    StringLit,
    OPERATOR,
    BOOL,
    RPAREN,
    LPAREN,
    RBRACE,
    LBRACE,
    RBRACKET,
    LBRACKET,
}

enum Mode {
    Friendly,
    Formal,
}

#[derive(ValueEnum, Debug, Clone)]
struct Args {
    name: String,
    count: u32,
    mode: Mode,
}

fn main() {
    let args = Args::parse();

    for _ in 0..args.count {
        match args.mode {
            Mode::Friendly => println!("Hello, {}!", args.name),
            Mode::Formal => println!("Good day, {}.", args.name),
        }
    }
}
