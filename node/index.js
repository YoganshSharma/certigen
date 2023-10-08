const Canvas = require('canvas')
const fs = require('fs')

async function createCertificate(image, fields) {
    const img = await Canvas.loadImage(image);
    const canvas = Canvas.createCanvas(img.width, img.height);
    const ctx = canvas.getContext('2d');
    ctx.drawImage(img, 0, 0);
    fields.forEach(field => {
        const height = Math.abs(field.ycord1 - field.ycord2);
        ctx.font = `${height}px Helvetica`;
        const fieldWidth =  Math.abs(field.xcord1 - field.xcord2);
        const startpos = (field.xcord1+field.xcord2 -Math.min(ctx.measureText(field.text).width, fieldWidth))/2;

        ctx.fillText(field.text,
            startpos,
            Math.max(field.ycord1, field.ycord2),
            fieldWidth,
        );
    });
    const buf = canvas.toBuffer();
    fs.writeFile("out.png", buf, () => { });
}

createCertificate("certi.png",
    [{ text: "Mr. Golla Meghanandh M Prabhash", xcord1: 1022.36, ycord1: 721.35, xcord2: 2049, ycord2: 801.3 },
    { text: "XYZ", xcord1: 403.6, ycord1: 832.9, xcord2: 1520, ycord2: 899.6 }]);


