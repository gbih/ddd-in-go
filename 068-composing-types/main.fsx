module Example068 =

    type CheckNumber = CheckNumber of int
    type CardNumber = CardNumber of string

    // 'OR' type
    type CardType =
    | Visa
    | Mastercard

    // 'AND' type (record)
    type CreditCardInfo = {
        CardType : CardType
        CardNumber : CardNumber
    }

    type PaymentMethod =
    | Cash
    | Check of CheckNumber
    | Card of CreditCardInfo

    type PaymentAmount = PaymentAmount of decimal
    type Currency = EUR | USD

    // top-level type
    type Payment = {
        Amount : PaymentAmount
        Currency: Currency
        Method: PaymentMethod
    }

    type ConvertPaymentCurrency = 
        Payment -> Currency -> Payment