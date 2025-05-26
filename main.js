const go = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});

function generate() {
    const useNum = document.getElementById("opt-num").checked;
    const useSym = document.getElementById("opt-symbol").checked;
    window.generatePassword(useNum, useSym);
}
