module Example084 =

    type Undefined = exn

    type CustomerInfo = Undefined 
    type ShippingAddress = Undefined 
    type BillingAddress = Undefined 
    type OrderLine = Undefined
    type BillingAmount = Undefined

    type Order = {
        CustomerInfo : CustomerInfo 
        ShippingAddress : ShippingAddress 
        BillingAddress : BillingAddress 
        OrderLines : OrderLine list 
        AmountToBill : BillingAmount
    }