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
