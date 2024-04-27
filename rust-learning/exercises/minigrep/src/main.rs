use minigrep::Config;
use std::{env, process};

fn main() {
    // let args: Vec<String> = env::args().collect();
    // dbg!(&args);

    let config = Config::build(env::args()).unwrap_or_else(|err| {
        eprintln!("Problem parsing arguments: {err}");
        process::exit(1);
    });

    // println!("Searching for {}", config.query);
    // println!("In file {}", config.file_path);

    if let Err(e) = minigrep::run(config) {
        eprintln!("Application error: {e}");
        process::exit(1);
    }
}
// Error message that crashed lsp
// Error ON_EXIT_CALLBACK_ERROR: "...es/start/rustaceanvim/lua/rustaceanvim/commands/init.lua:323: E5560: nvim_del_user_command m
// ust not be called in a lua loop callback"
