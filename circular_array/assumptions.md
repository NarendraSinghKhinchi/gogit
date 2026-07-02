Design a Circular Array

Constraints and assumptions

Is the circular array generic or does it store only integers?
Assume it is generic.

By "circular", do you mean a circular view of a fixed-size array or a circular buffer where new elements overwrite old ones?
Assume it is a circular view of a fixed-size array. Elements are never overwritten.

Should rotation be implemented?
Yes, support rotating the array in both clockwise and counter-clockwise directions.

Should rotation physically move the elements?
No, rotation should be efficient and should not require rearranging the underlying array.

Can the size of the array change after creation?
No, assume the size is fixed after initialization.

Should the array support random access by index?
Yes, it should support get(index) and set(index).

Should iteration reflect the rotated order?
Yes.

Can we assume inputs are valid, or do we need to handle invalid indices and rotations?
Assume inputs are valid.