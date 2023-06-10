# ET_Playground

This is a work-in-progress playground used for testing out ideas within Element Theory created from scratch using Raylib. Current features include:
- A six-element system (|A|, |I|, |U|, |S|, |H| and |L|, where |S| represents the stop element). Each element is added to the field by pressing its corresponding key on the keyboard ('a' for the |A| element, 'l' for the |L| element, and so on). Right click to remove an element.
- Headed and non-headed elements. A headed element is added by holding shift while adding it to the field.
- Combining elements to form segments. Drag and drop elements into one another to form a segment. Hold shift to gain access to an element inside (for instance if you need to move it out of the segment or remove it from the field entirely).
- Ctrl-click an element to create a segment consisting of only one element.
- You may create a simplex segment and then subsequently delete its only element to create an empty segment.
- Right click on a segment to delete it. All of its elements will remain on the field and will have to be removed separately if necessary. This is intended behavior.

![demo](demo/et_pg_demo.gif)