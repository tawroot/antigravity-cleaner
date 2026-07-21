from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]


def test_native_shell_script_exists_and_is_selected_by_unix_installer():
    script = ROOT / "Antigravity.sh"
    assert script.exists(), "Antigravity.sh should provide native macOS/Linux implementation"
    install = (ROOT / "install.sh").read_text()
    assert "Antigravity.sh" in install
    assert "pwsh" not in install.lower()


def test_native_shell_covers_main_ps1_menu_modules():
    script = (ROOT / "Antigravity.sh").read_text()
    for token in [
        "System Cleaner",
        "Session Manager",
        "Network Optimizer",
        "System Analysis",
        "Region Inspector",
        "backup_browser_profile",
        "restore_profile",
        "network_reset",
    ]:
        assert token in script


def test_native_shell_uses_macos_and_linux_paths_not_windows_globs():
    script = (ROOT / "Antigravity.sh").read_text()
    assert "Library/Application Support/Google/Chrome" in script
    assert ".config/google-chrome" in script
    assert "\\*" not in script
    assert "powershell" not in script.lower()
    assert "pwsh" not in script.lower()


if __name__ == "__main__":
    tests = [
        test_native_shell_script_exists_and_is_selected_by_unix_installer,
        test_native_shell_covers_main_ps1_menu_modules,
        test_native_shell_uses_macos_and_linux_paths_not_windows_globs,
    ]
    for test in tests:
        test()
        print(f"ok - {test.__name__}")
