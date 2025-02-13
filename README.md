# Codebase Structure

```
📂 .
│   📂 .github
│       📂 workflows
│           📄 generate-tree.yml
│   📄 .gitignore
│   📂 1.0.array
│   │   📂 1.0.sum-array
│   │   │   📄 readme.md
│   │       📄 sum-array.go
│   │   📂 1.1.sequential-search
│   │   │   📄 readme.md
│   │       📄 sequential-search.go
│   │   📂 1.2.binary-search
│   │   │   📄 binary-search.go
│   │       📄 readme.md
│   │   📂 1.3.finding-the-largest-sum-subarray
│   │   │   📄 largest-sum.go
│   │       📄 readme.md
│   │   📂 1.4.rotate-an-array-by-k-position
│   │   │   📄 readme.md
│   │       📄 rotate-k.go
│   │   📂 1.5.array-waveform
│   │   │   📄 array-waveform.go
│   │       📄 readme.md
│   │   📂 1.6.index-array
│   │   │   📄 index-array.go
│   │       📄 readme.md
│   │   📂 1.7.sorting-from-1-to-n
│   │   │   📄 readme.md
│   │       📄 sort.go
│   │   📂 1.8.smallest-positive-missing-number
│   │       📄 missing-number.go
│   │   📂 1.9.contains-duplicate
│   │   │   📄 contains-duplicate.go
│   │       📄 readme.md
│   │   📂 2.0.random
│   │       📄 random.go
│   │   📂 2.1.valid-anagram
│   │   │   📄 readme.md
│   │       📄 valid-anagrams.go
│   │   📂 2.2.two-sum
│   │   │   📄 readme.md
│   │       📄 two-sum.go
│   │   📂 2.3.group-anagrams
│   │   │   📄 group-anagrams.go
│   │       📄 readme.md
│   │   📂 2.4.top-k-frequent-element
│   │   │   📄 k-frequent.go
│   │       📄 readme.md
│   │   📂 2.5.encode-and-decode-strings
│   │   │   📄 encode-decode.go
│   │       📄 readme.md
│   │   📂 2.6.is-a-leap-year
│   │   │   📄 leap-year.go
│   │       📄 readme.md
│   │   📂 2.7.isogram
│   │   │   📄 isogram.go
│   │       📄 readme.md
│   │   📂 2.8.scrabble
│   │   │   📄 readme.md
│   │       📄 scrabble.go
│   │   📂 2.9.trinary
│   │   │   📄 readme.md
│   │       📄 trinary.go
│   │   📂 3.0.encode-length
│   │   │   📄 encode-length.go
│   │       📄 readme.md
│   │   📂 3.1.product-except-self
│   │       📄 product.go
│   │   📂 3.2.valid-sudoku
│   │       📄 valid-sudoku.go
│   │   📂 3.3.max-area
│   │   │   📄 max-area.go
│   │       📄 readme.md
│   │   📂 3.4.valid-mountian-array
│   │   │   📄 readme.md
│   │       📄 valid-mountain.go
│   │   📂 3.5.boats-required
│   │   │   📄 boats-required.go
│   │       📄 readme.md
│   │   📂 3.6.move-zeroes
│   │   │   📄 move-zeroes.go
│   │       📄 readme.md
│       📂 3.7.longest-substring
│       │   📄 longest-substring.go
│           📄 readme.md
│   📂 1.1.linked-list
│   │   📂 doubly-linked-list
│   │   │   📂 1.0.create-a-vowel-double-linked-list
│   │   │   │   📄 main.go
│   │   │       📄 readme.md
│   │   │   📄 main.go
│   │       📄 readme.md
│   │   📄 readme.md
│       📂 single-linked-list
│       │   📂 1.0.create-a-vowel-linked-list
│       │   │   📄 main.go
│       │       📄 readme.md
│           📄 main.go
│   📂 1.2.stacks
│   │   📂 1.0.vowel-implementation-of-stack
│   │       📄 main.go
│   │   📄 main.go
│       📄 readme.md
│   📂 1.3.queue
│   │   📂 1.0.implement-queue-with-channel
│   │       📄 main.go
│   │   📄 main.go
│       📄 readme.md
│   📂 1.4.binary-tree
│   │   📂 array-implementation-of-binary-tree
│   │       📄 main.go
│       📄 main.go
│   📂 1.5.concurrency
│   │   📂 1.0.channel-example
│   │       📄 main.go
│   │   📂 1.1.fan-in
│   │       📄 main.go
│   │   📂 1.2.fan-out
│   │       📄 main.go
│   │   📂 1.3.worker-pool
│   │       📄 main.go
│       📂 1.4.distributed-search-engine
│           📄 main.go
│   📂 1.6.interface
│   │   📄 main.go
│       📄 readme.md
│   📂 1.7.struct
│       📄 readme.md
│   📂 1.8.testing
│       📄 test_main.go
│   📂 1.9.design-patterns
│   │   📂 1.0.creational
│   │   │   📂 1.0.singleton
│   │   │   │   📄 readme.md
│   │   │   │   📄 singleton.go
│   │   │       📄 singleton_test.go
│   │   │   📂 2.0.builder
│   │   │   │   📄 builder.go
│   │   │   │   📄 readme.md
│   │   │       📄 test_builder.go
│   │   │   📂 3.0.factory
│   │   │   │   📄 factory.go
│   │   │   │   📄 factory_test.go
│   │   │       📄 readme.md
│   │   │   📂 4.0.abstract-factory
│   │   │   │   📄 car-factory.go
│   │   │   │   📄 car.go
│   │   │   │   📄 cruise-motorbike.go
│   │   │   │   📄 family-car.go
│   │   │   │   📄 luxury-car.go
│   │   │   │   📄 motorbike-factory.go
│   │   │   │   📄 motorbike.go
│   │   │   │   📄 readme.md
│   │   │   │   📄 sport-motorbike.go
│   │   │   │   📄 vechicle-factory_test.go
│   │   │   │   📄 vehicle-factory.go
│   │   │       📄 vehicle.go
│   │   │   📂 5.0.prototype
│   │   │   │   📄 prototype.go
│   │   │   │   📄 prototype_test.go
│   │   │       📄 readme.md
│   │       📄 readme.md
│       📂 2.0.structural
│       │   📂 1.0.composition
│       │   │   📂 1.0.composition
│       │   │   │   📄 composition.go
│       │   │   │   📄 composition_test.go
│       │   │       📄 readme.md
│       │       📂 1.1.binary-tree-composition
│       │       │   📄 binary-tree.go
│       │       │   📄 binary-tree_test.go
│       │           📄 readme.md
│       │   📂 2.0.adapter
│       │   │   📄 adapter.go
│       │   │   📄 adapter_test.go
│       │       📄 readme.md
│           📂 3.0.bridge
│           │   📄 bridge.go
│           │   📄 bridge_test.go
│               📄 readme.md
│   📄 README.md
│   📂 generate
│       📄 generate.go
│   📄 go.mod
│   📄 keys.go
    📄 main.go
```
