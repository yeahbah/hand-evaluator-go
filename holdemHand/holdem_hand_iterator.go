package holdemHand

// slow
func cardsRange(mask uint64) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 51; i >= 0; i-- {
			if (uint64(1)<<i)&mask != 0 {
				channel <- CardTable[i]
			}
		}
		close(channel)
	}()

	return channel
}

func cardsRange2(mask uint64, callback func(string)) {
	for i := 51; i >= 0; i-- {
		if (uint64(1)<<i)&mask != 0 {
			callback(CardTable[i])
		}

	}
}

// slow
func HandsRange(numCards int) <-chan uint64 {
	channel := make(chan uint64)

	switch numCards {
	case 7:
		go func() {
			for a := 0; a < CardsMasksTableSize-6; a++ {
				_card1 := CardMasksTable[a]
				for b := a + 1; b < CardsMasksTableSize-5; b++ {
					_n2 := _card1 | CardMasksTable[b]
					for c := b + 1; c < CardsMasksTableSize-4; c++ {
						_n3 := _n2 | CardMasksTable[c]
						for d := c + 1; d < CardsMasksTableSize-3; d++ {
							_n4 := _n3 | CardMasksTable[d]
							for e := d + 1; e < CardsMasksTableSize-2; e++ {
								_n5 := _n4 | CardMasksTable[e]
								for f := e + 1; f < CardsMasksTableSize-1; f++ {
									_n6 := _n5 | CardMasksTable[f]
									for g := f + 1; g < CardsMasksTableSize; g++ {
										channel <- _n6 | CardMasksTable[g]
									}
								}
							}
						}
					}
				}
			}
			close(channel)
		}()
	case 6:
		go func() {
			for a := 0; a < CardsMasksTableSize-5; a++ {
				_card1 := CardMasksTable[a]
				for b := a + 1; b < CardsMasksTableSize-4; b++ {
					_n2 := _card1 | CardMasksTable[b]
					for c := b + 1; c < CardsMasksTableSize-3; c++ {
						_n3 := _n2 | CardMasksTable[c]
						for d := c + 1; d < CardsMasksTableSize-2; d++ {
							_n4 := _n3 | CardMasksTable[d]
							for e := d + 1; e < CardsMasksTableSize-1; e++ {
								_n5 := _n4 | CardMasksTable[e]
								for f := e + 1; f < CardsMasksTableSize; f++ {
									channel <- _n5 | CardMasksTable[f]
								}
							}
						}
					}
				}
			}
			close(channel)
		}()
	case 5:
		go func() {
			for a := 0; a < CardsMasksTableSize-4; a++ {
				_card1 := CardMasksTable[a]
				for b := a + 1; b < CardsMasksTableSize-3; b++ {
					_n2 := _card1 | CardMasksTable[b]
					for c := b + 1; c < CardsMasksTableSize-2; c++ {
						_n3 := _n2 | CardMasksTable[c]
						for d := c + 1; d < CardsMasksTableSize-1; d++ {
							_n4 := _n3 | CardMasksTable[d]
							for e := d + 1; e < CardsMasksTableSize; e++ {
								channel <- _n4 | CardMasksTable[e]
							}
						}
					}
				}
			}
			close(channel)
		}()
	case 4:
		go func() {
			for a := 0; a < CardsMasksTableSize-3; a++ {
				_card1 := CardMasksTable[a]
				for b := a + 1; b < CardsMasksTableSize-2; b++ {
					_n2 := _card1 | CardMasksTable[b]
					for c := b + 1; c < CardsMasksTableSize-1; c++ {
						_n3 := _n2 | CardMasksTable[c]
						for d := c + 1; d < CardsMasksTableSize; d++ {
							channel <- _n3 | CardMasksTable[d]
						}
					}

				}
			}
			close(channel)
		}()
	case 3:
		go func() {
			for a := 0; a < CardsMasksTableSize-2; a++ {
				_card1 := CardMasksTable[a]
				for b := a + 1; b < CardsMasksTableSize-1; b++ {
					_n2 := _card1 | CardMasksTable[b]
					for c := b + 1; c < CardsMasksTableSize; c++ {
						channel <- _n2 | CardMasksTable[c]
					}
				}
			}
			close(channel)
		}()
	case 2:
		go func() {
			for a := 0; a < TwoCardMaskTableSize; a++ {
				channel <- TwoCardMaskTable[a]
			}
			close(channel)
		}()
	case 1:
		go func() {
			for a := 0; a < CardsMasksTableSize; a++ {
				channel <- CardMasksTable[a]
			}
			close(channel)
		}()
	default:
		channel <- 0
		close(channel)
	}

	return channel
}

func HandsRange2(numCards int, callback func(uint64)) {

	switch numCards {
	case 7:
		for a := 0; a < CardsMasksTableSize-6; a++ {
			_card1 := CardMasksTable[a]
			for b := a + 1; b < CardsMasksTableSize-5; b++ {
				_n2 := _card1 | CardMasksTable[b]
				for c := b + 1; c < CardsMasksTableSize-4; c++ {
					_n3 := _n2 | CardMasksTable[c]
					for d := c + 1; d < CardsMasksTableSize-3; d++ {
						_n4 := _n3 | CardMasksTable[d]
						for e := d + 1; e < CardsMasksTableSize-2; e++ {
							_n5 := _n4 | CardMasksTable[e]
							for f := e + 1; f < CardsMasksTableSize-1; f++ {
								_n6 := _n5 | CardMasksTable[f]
								for g := f + 1; g < CardsMasksTableSize; g++ {
									callback(_n6 | CardMasksTable[g])
								}
							}
						}
					}
				}
			}
		}

	case 6:
		for a := 0; a < CardsMasksTableSize-5; a++ {
			_card1 := CardMasksTable[a]
			for b := a + 1; b < CardsMasksTableSize-4; b++ {
				_n2 := _card1 | CardMasksTable[b]
				for c := b + 1; c < CardsMasksTableSize-3; c++ {
					_n3 := _n2 | CardMasksTable[c]
					for d := c + 1; d < CardsMasksTableSize-2; d++ {
						_n4 := _n3 | CardMasksTable[d]
						for e := d + 1; e < CardsMasksTableSize-1; e++ {
							_n5 := _n4 | CardMasksTable[e]
							for f := e + 1; f < CardsMasksTableSize; f++ {
								callback(_n5 | CardMasksTable[f])
							}
						}
					}
				}
			}
		}

	case 5:

		for a := 0; a < CardsMasksTableSize-4; a++ {
			_card1 := CardMasksTable[a]
			for b := a + 1; b < CardsMasksTableSize-3; b++ {
				_n2 := _card1 | CardMasksTable[b]
				for c := b + 1; c < CardsMasksTableSize-2; c++ {
					_n3 := _n2 | CardMasksTable[c]
					for d := c + 1; d < CardsMasksTableSize-1; d++ {
						_n4 := _n3 | CardMasksTable[d]
						for e := d + 1; e < CardsMasksTableSize; e++ {
							callback(_n4 | CardMasksTable[e])
						}
					}
				}
			}
		}

	case 4:

		for a := 0; a < CardsMasksTableSize-3; a++ {
			_card1 := CardMasksTable[a]
			for b := a + 1; b < CardsMasksTableSize-2; b++ {
				_n2 := _card1 | CardMasksTable[b]
				for c := b + 1; c < CardsMasksTableSize-1; c++ {
					_n3 := _n2 | CardMasksTable[c]
					for d := c + 1; d < CardsMasksTableSize; d++ {
						callback(_n3 | CardMasksTable[d])
					}
				}

			}
		}

	case 3:

		for a := 0; a < CardsMasksTableSize-2; a++ {
			_card1 := CardMasksTable[a]
			for b := a + 1; b < CardsMasksTableSize-1; b++ {
				_n2 := _card1 | CardMasksTable[b]
				for c := b + 1; c < CardsMasksTableSize; c++ {
					callback(_n2 | CardMasksTable[c])
				}
			}
		}

	case 2:

		for a := 0; a < TwoCardMaskTableSize; a++ {
			callback(TwoCardMaskTable[a])
		}

	case 1:

		for a := 0; a < CardsMasksTableSize; a++ {
			callback(CardMasksTable[a])
		}

	default:
		return

	}

}
