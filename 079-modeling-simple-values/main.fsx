module Example079 =

    type CustomerId =
    | CustomerId of int

    type WidgetCode = WidgetCode of string
    type UnitQuantity = UnitQuantity of int
    type KilogramQuantity = KilogramQuantity of decimal


    // Deconstruct a single case union
    // construct
    let customerId = CustomerId 42

    // deconstruct
    let (CustomerId innerValue) = customerId
    printfn "%i" innerValue

    // deconstruct directly in the parameter of a function definition
    let processCustomerId (CustomerId innerValue) =
        printfn "innerValue is %i" innerValue

    processCustomerId customerId