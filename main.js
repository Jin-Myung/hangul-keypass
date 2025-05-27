const go = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});

let cooldown = false;

function generate() {
    if (cooldown) return;

    cooldown = true;
    const btn = document.getElementById("generate-btn");
    btn.disabled = true;
    const originalText = btn.innerText;
    btn.innerText = "생성 중...";

    const useNum = document.getElementById("opt-num").checked;
    const useSym = document.getElementById("opt-symbol").checked;
    window.generatePassword(useNum, useSym);

    setTimeout(() => {
        cooldown = false;
        btn.disabled = false;
        btn.innerText = originalText;
    }, 1200); // 1.2초 쿨다운
}
