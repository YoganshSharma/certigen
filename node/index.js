const Canvas = require('canvas')
const fs = require('fs')

async function createCertificate(image, fields) {
    const img = await Canvas.loadImage(image);
    const canvas = Canvas.createCanvas(img.width, img.height);
    const ctx = canvas.getContext('2d');
    ctx.drawImage(img, 0, 0);
    fields.forEach(field => {
        const height = Math.abs(field.ycord1-field.ycord2);
        ctx.font = `${height}px Helvetica`;
        ctx.fillText(field.text,
            Math.min(field.xcord1, field.xcord1),
            Math.max(field.ycord1, field.ycord2),
            Math.abs(field.xcord1- field.xcord2),
        );
    });
    const buf = canvas.toBuffer();
    fs.writeFile("out.png", buf, () => { });
}

createCertificate("certi.png", [{ text: "Mr. Golla Meghanandh M Prabhash", xcord1: 1022.36, ycord1: 721.35, xcord2: 2049, ycord2: 801.3 }]);


