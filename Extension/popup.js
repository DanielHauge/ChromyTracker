productList = ['Electronics Watch','House wear Items','Kids wear','Women Fashion'];


$(function(){

    var ul = document.createElement('ul');
    ul.setAttribute('id', 'proList');

    var t, tt;
    
    
    
    document.getElementById('renderList').appendChild(ul);
    productList.forEach(renderProductList);

    function renderProductList(element, index, arr) {
        var li = document.createElement('li');
        var button = document.createElement("input");
        button.type = "checkbox";
        button.setAttribute('id', element+'id');
        
        button.innerHTML = element;
        li.setAttribute('class','checkbox');
        li.appendChild(button)

        ul.appendChild(li);

        t = document.createTextNode(element);

        li.innerHTML=li.innerHTML+element;
    };

    
    
    

    $('#spendAmount').click(function(){
        chrome.storage.sync.get(['total','limit'], function(budget){
            var newTotal = 0;
            if (budget.total){
                newTotal += parseInt(budget.total);
            }


            var amount = $('#amount').val();
            if (amount){
                newTotal += parseInt(amount);
            }


            chrome.storage.sync.set({'total':newTotal}, function(){
                if (amount && newTotal>=budget.limit){
                    var notifOptions = {
                        type: 'basic',
                        iconUrl: 'icon48.png',
                        title: 'Limit reached',
                        message: 'Uh oh!'
                    };

                    chrome.notifications.create('limitNotif', notifOptions);
                }
            });

            $('#total').text(newTotal);
            $('#amount').val('');

        });
    });
});






    

