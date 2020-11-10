module Example067 =

    type OrderQuantity =
    | UnitQuantity of int
    | KilogramQuantity of float

    let anOrderQtyInUnits = UnitQuantity 10 
    let anOrderQtyInKg = KilogramQuantity 2.5

    let printQuantity aOrderQty = 
        match aOrderQty with
        | UnitQuantity uQty ->
            printfn "%i units" uQty
        | KilogramQuantity kgQty ->
            printfn "%g kg" kgQty

    printQuantity anOrderQtyInUnits // "10 units" 
    printQuantity anOrderQtyInKg // "2.5 kg"