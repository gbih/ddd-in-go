# pg.32

Original pseudo-code:
```
bounded context: Order-Taking
data Order =
    CustomerInfo
    AND ShippingAddress
    AND BillingAddress
    AND list of OrderLines
    AND AmountToBill
data OrderLine =
    Product
    AND Quantity
    AND Price
data CustomerInfo = ???
data BillingAddress = ??? // don't know yet
```
