package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestDeleteByPath(t *testing.T) {
	mp := map[string]any{
		"products": map[string]any{
			"desk": map[string]int{
				"price": 100,
			},
		},
	}
	assert.True(t, maputil.DeleteByPath(mp, "products.desk"))
	assert.Equal(t, map[string]any{
		"products": map[string]any{},
	}, mp)

	mp = map[string]any{
		"products": map[string]any{
			"desk": map[string]int{
				"price": 100,
			},
		},
	}
	assert.True(t, maputil.DeleteByPath(mp, "products.desk.price"))
	assert.False(t, maputil.DeleteByPath(mp, "products.desk.quantity"))
	assert.Equal(t, map[string]any{
		"products": map[string]any{
			"desk": map[string]int{},
		},
	}, mp)

	mp = map[string]any{
		"products": map[string]any{
			"desk": map[string]int{
				"price": 100,
			},
		},
	}
	assert.False(t, maputil.DeleteByPath(mp, "products.chair.price"))
	assert.Equal(t, map[string]any{
		"products": map[string]any{
			"desk": map[string]int{
				"price": 100,
			},
		},
	}, mp)

	mp = map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": map[string]int{
					"original": 100,
					"taxes":    120,
				},
			},
		},
	}
	assert.True(t, maputil.DeleteByPath(mp, "products.desk.price.taxes"))
	assert.Equal(t, map[string]any{
		"products": map[string]any{
			"desk": map[string]any{
				"price": map[string]int{
					"original": 100,
				},
			},
		},
	}, mp)

	mp = map[string]any{
		"developers": []map[string]any{
			{
				"name": "John",
				"lang": []string{
					"Golang",
					"PHP",
				},
			},
			{
				"name": "Krishan",
				"lang": "Golang",
			},
		},
	}
	assert.True(t, maputil.DeleteByPath(mp, "developers.*.lang"))
	assert.Equal(t, map[string]any{
		"developers": []map[string]any{
			{
				"name": "John",
			},
			{
				"name": "Krishan",
			},
		},
	}, mp)

	mp = map[string]any{
		"developers": []map[string]any{
			{
				"name": "John",
				"lang": []string{
					"Golang",
					"PHP",
				},
			},
			{
				"name": "Krishan",
				"lang": []string{
					"C",
					"Golang",
				},
			},
		},
	}
	assert.True(t, maputil.DeleteByPath(mp, "developers.*.lang.1"))
	assert.True(t, maputil.DeleteByPath(mp, "developers.*.name"))
	assert.Equal(t, map[string]any{
		"developers": []map[string]any{
			{
				"lang": []string{
					"Golang",
				},
			},
			{
				"lang": []string{
					"C",
				},
			},
		},
	}, mp)

	// Only works on first level keys
	mp = map[string]any{
		"joe@example.com": "Joe",
		"jane@localhost":  "Jane",
	}
	assert.True(t, maputil.DeleteByPath(mp, "joe@example.com"))
	assert.Equal(t, map[string]any{
		"jane@localhost": "Jane",
	}, mp)

	// Doesn't remove nested keys
	mp = map[string]any{
		"emails": map[string]string{
			"joe@example.com": "Joe",
			"jane@localhost":  "Jane",
		},
	}
	assert.False(t, maputil.DeleteByPath(mp, "emails.joe@example.com"))
	assert.Equal(t, map[string]any{
		"emails": map[string]string{
			"joe@example.com": "Joe",
			"jane@localhost":  "Jane",
		},
	}, mp)

	mp = map[string]any{
		"developers": []map[string]string{
			{
				"name": "John",
			},
			{
				"name": "Krishan",
			},
		},
	}
	assert.True(t, maputil.DeleteByPath(mp, "developers.*.name"))
	assert.Equal(t, map[string]any{
		"developers": []map[string]string{},
	}, mp)

	// Test nil value
	mp = map[string]any{
		"shop": map[string]any{
			"cart": map[any]any{
				150:   0,
				"foo": "bar",
			},
		},
	}
	assert.True(t, maputil.DeleteByPath(mp, "shop.cart.150"))
	assert.True(t, maputil.DeleteByPath(mp, "shop.cart.100"))
	assert.True(t, maputil.DeleteByPath(mp, "shop.cart.foo"))
	assert.Equal(t, map[string]any{
		"shop": map[string]any{
			"cart": map[any]any{},
		},
	}, mp)

	mp = map[string]any{
		"developers": []map[string]any{
			{
				"lang": []string{
					"Golang",
					"PHP",
				},
			},
		},
	}
	assert.True(t, maputil.DeleteByPath(mp, "developers.0.lang.0"))
	assert.Equal(t, map[string]any{
		"developers": []map[string]any{
			{
				"lang": []string{
					"PHP",
				},
			},
		},
	}, mp)

	mp = map[string]any{
		"names": []string{"Bowen", "Krishan"},
	}
	assert.False(t, maputil.DeleteByPath(mp, "names.*.foo"))
	assert.False(t, maputil.DeleteByPath(mp, "names.3"))
	assert.False(t, maputil.DeleteByPath(mp, "names.*"))
	assert.Eq(t, map[string]any{
		"names": []string{},
	}, mp)

}
