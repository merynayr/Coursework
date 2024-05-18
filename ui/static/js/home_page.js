function allBlock(){
    var e1 = document.getElementById(1);
    var e2 = document.getElementById(2);
    var e3 = document.getElementById(3);
    var e4 = document.getElementById(4);
    var e5 = document.getElementById(5);
    e1.style.display = 'block';
    e2.style.display = 'none';
    e3.style.display = 'none';
    e4.style.display = 'none';
    e5.style.display = 'none';
}
window.onload = allBlock;
function toggle_visibility(id) {
    var e1 = document.getElementById(1);
    var e2 = document.getElementById(2);
    var e3 = document.getElementById(3);
    var e4 = document.getElementById(4);
    var e5 = document.getElementById(5);
    switch(id){
        case 1:
            if(e1.style.display != 'block')
                e1.style.display = 'block';
            if(e2.style.display != 'none')
                e2.style.display = 'none';
            if(e3.style.display != 'none')
                e3.style.display = 'none';
            if(e4.style.display != 'none')
                e4.style.display = 'none';
            if(e5.style.display != 'none')
                e5.style.display = 'none';
            break;
        case 2:
            if(e1.style.display != 'none')
                e1.style.display = 'none';
            if(e2.style.display != 'block')
                e2.style.display = 'block';
            if(e3.style.display != 'none')
                e3.style.display = 'none';
            if(e4.style.display != 'none')
                e4.style.display = 'none';
            if(e5.style.display != 'none')
                e5.style.display = 'none';
            break;
         case 3:
            if(e1.style.display != 'none')
                e1.style.display = 'none';
            if(e2.style.display != 'none')
                e2.style.display = 'none';
            if(e3.style.display != 'block')
                e3.style.display = 'block';
            if(e4.style.display != 'none')
                e4.style.display = 'none';
            if(e5.style.display != 'none')
                e5.style.display = 'none';
            break;
         case 4:
            if(e1.style.display != 'none')
                e1.style.display = 'none';
            if(e2.style.display != 'none')
                e2.style.display = 'none';
            if(e3.style.display != 'none')
                e3.style.display = 'none';
            if(e4.style.display != 'block')
                e4.style.display = 'block';
            if(e5.style.display != 'none')
                e5.style.display = 'none';
            break;
         case 5:
            if(e1.style.display != 'none')
                e1.style.display = 'none';
            if(e2.style.display != 'none')
                e2.style.display = 'none';
            if(e3.style.display != 'none')
                e3.style.display = 'none';
            if(e4.style.display != 'none')
                e4.style.display = 'none';
            if(e5.style.display != 'block')
                e5.style.display = 'block';
            break;
    }

}

function formatDate(date) {
    var date = new Date(date);
    var yyyy = date.getFullYear().toString();
    var mm = (date.getMonth() + 1).toString();
    var dd = date.getDate().toString();

    var mmChars = mm.split('');
    var ddChars = dd.split('');

    return ((ddChars[1] ? dd : "0" + ddChars[0])  + '.' + (mmChars[1] ? mm : "0" + mmChars[0]) + '.' + yyyy);
}

//  Добавляет ведущие нули в число, например 9 = 09
Number.prototype.pad = function(size) {
    var s = String(this);
    while (s.length < (size || 2)) {s = "0" + s;}
    return s;
}