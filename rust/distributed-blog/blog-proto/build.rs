use std::{env, fs};
use std::path::Path;

fn main(){
    let current_dir = env::current_dir().unwrap();
    let proto_path = Path::new(&current_dir).join("proto");
    let mut proto_files = vec![];
    for entry in fs::read_dir(&proto_path).unwrap() {
        let entry = entry.unwrap();
        let md = entry.metadata().unwrap();
        if md.is_file() && entry.path().extension().unwrap() == "proto" {
            proto_files.push(entry.path().as_os_str().to_os_string())
        }
    }

    tonic_build::configure()
        .out_dir("src")
        .build_client(true)
        .build_server(true)
        .compile(
            proto_files.as_slice(),
            &[&proto_path],
        )
        .unwrap();
}