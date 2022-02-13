const c = document.getElementById("c");

const step = document.getElementById("step");
const k = document.getElementById("k");
const u = document.getElementById("u");
const t = document.getElementById("t");

const source = new EventSource("/sse/lj");

var L, iL;

source.addEventListener("reset", (event) => {
    const data = JSON.parse(event.data);
    L = data.L;
    iL = 1 / L;
});

source.addEventListener("update", (event) => {
    const data = JSON.parse(event.data);
    step.innerText = data.t;
    k.innerText = data.K;
    u.innerText = data.U;
    t.innerText = data.T;
    draw(data.pos);
});

function draw(pos) {
    let ctx = c.getContext("2d");
    let w = c.clientWidth;
    let h = c.clientHeight;

    ctx.fillStyle = "#000";
    ctx.fillRect(0, 0, w, h);

    pos.forEach((p) => {
        let x = (p[0]*iL + 0.5) * w;
        let y = (p[1]*iL + 0.5) * h;
        let z = parseInt((p[2]*iL + 0.5) * 200 + 55);
        ctx.fillStyle = `rgba(0, 0, ${z}, 1)`;
        ctx.beginPath();
        ctx.arc(x, y, 10, 0, 2*Math.PI);
        ctx.fill();
    })
}
