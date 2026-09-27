//! Recursively searches for a file by name in a directory.
//! Demonstrates Rust's safety, error handling, and performance.

use std::env;
use std::fs;
use std::path::Path;
use chrono::Local;

/// Prints system information such as OS, architecture, CPU count, and current time.
fn print_system_info() {
    println!("--- System Info ---");
    println!("OS: {}", env::consts::OS);
    println!("Architecture: {}", env::consts::ARCH);
    println!("Number of CPUs: {}", num_cpus::get());
    println!("Current Time: {}", Local::now().format("%a, %d %b %Y %H:%M:%S %Z"));
    println!("-------------------");
}

/// Recursively searches for files with a given name in a directory.
/// # Arguments
/// * `dir` - The directory to search in.
/// * `target` - The filename to search for.
/// * `results` - A mutable reference to a vector to store results.
/// # Returns
/// `std::io::Result<()>` - Ok if successful, Err otherwise.
fn find_file_recursively(dir: &Path, target: &str, results: &mut Vec<String>) -> std::io::Result<()> {
    if dir.is_dir() {
        for entry in fs::read_dir(dir)? {
            let entry = entry?;
            let path = entry.path();
            if path.is_dir() {
                find_file_recursively(&path, target, results)?;
            } else if let Some(name) = path.file_name() {
                if name.to_string_lossy().contains(target) {
                    results.push(path.display().to_string());
                }
            }
        }
    }
    Ok(())
}

fn main() {
    println!("Hello from Rust!");

    print_system_info();

    println!("\nRecursively searching for files named 'Cargo.toml'...");
    let mut results = Vec::new();
    if let Err(err) = find_file_recursively(Path::new("."), "Cargo.toml", &mut results) {
        eprintln!("Error: {}", err);
    } else {
        println!("Found {} files:", results.len());
        for file in results {
            println!("  - {}", file);
        }
    }

    println!("\nExiting...");
}