use std::fs::File;
use std::io::prelude::*;
use std::path::Path;

use prost_build;

const PROTO_DIR: &str = "../proto";
const RUST_DIR: &str = "../rust";

fn main() {
    let paths_iter = std::fs::read_dir(PROTO_DIR)
        .unwrap()
        .map(|entry| entry.unwrap().file_name().to_owned())
        .filter(|path| path.to_str().unwrap().ends_with(".proto"));
    let paths = Vec::from_iter(paths_iter);

    std::env::set_var("OUT_DIR", Path::new(RUST_DIR).join("src").to_str().unwrap());

    prost_build::compile_protos(&paths, &[PROTO_DIR])
        .unwrap();

    let mut lib_file = File::create(Path::new(RUST_DIR).join("src").join("lib.rs")).unwrap();
    for path in &paths {
        let mod_name = Path::new(path)
            .components()
            .last()
            .unwrap()
            .as_os_str()
            .to_str()
            .unwrap()
            .strip_suffix(".proto")
            .unwrap()
            .to_owned();
        lib_file.write_all(format!("pub mod {};\n", mod_name).as_bytes())
            .unwrap();
    }
}
