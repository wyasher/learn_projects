use std::{env, process};
use grep::{Config, run};
// cargo run -- to sample.txt
fn main() {
    // args_os 对非Unicode字符比较友好
    // args 无法处理非Unicode字符
    let config= Config::build(env::args())
        .unwrap_or_else(|err| {
            eprintln!("Problem parsing arguments: {err}");
            process::exit(1);
        });
    if let Err(e) = run(config){
        eprintln!("application error : {e}");
        process::exit(1);
    }
}
