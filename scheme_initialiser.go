package sumUp

func GetSchemes() []*Scheme {
	amexScheme := NewScheme("American Express",
		[]bound{
			{
				lower: 34,
				upper: 34,
			},
			{
				lower: 37,
				upper: 37,
			},
		},
		[]bound{
			{
				lower: 15,
				upper: 15,
			},
		},
	)

	jcbScheme := NewScheme("JCB",
		[]bound{
			{
				lower: 3528,
				upper: 3589,
			},
		},
		[]bound{
			{
				lower: 16,
				upper: 19,
			},
		},
	)

	maestroScheme := NewScheme("Maestro",
		[]bound{
			{
				lower: 50,
				upper: 50,
			},
			{
				lower: 56,
				upper: 58,
			},
			{
				lower: 6,
				upper: 6,
			},
		},
		[]bound{
			{
				lower: 12,
				upper: 19,
			},
		},
	)

	visaScheme := NewScheme("Visa",
		[]bound{
			{
				lower: 4,
				upper: 4,
			},
		},
		[]bound{
			{
				lower: 13,
				upper: 13,
			},
			{
				lower: 16,
				upper: 16,
			},
			{
				lower: 19,
				upper: 19,
			},
		},
	)

	mastercardScheme := NewScheme("MasterCard",
		[]bound{
			{
				lower: 2221,
				upper: 2720,
			},
			{
				lower: 51,
				upper: 55,
			},
		},
		[]bound{
			{
				lower: 16,
				upper: 16,
			},
		},
	)

	return []*Scheme{amexScheme, jcbScheme, maestroScheme, visaScheme, mastercardScheme}
}
