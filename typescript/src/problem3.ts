function getInput3(): string {
    return `..##.......
    #...#...#..
    .#....#..#.
    ..#.#...#.#
    .#...##..#.
    ..#.##.....
    .#.#.#....#
    .#........#
    #.##...#...
    #...##....#
    .#..#...#.#`;
}

enum Thing {
    Tree,
    Snow,
}

const things = getInput3().split('\n').map(x => x.split("").map(x => x === "." ? Thing.Snow : Thing.Tree));

const colLength = things[0].length;
let treeCount = 0.

things.forEach((row, i) => {
    if (row[i * 3 % colLength] === Thing.Tree) {
        treeCount++;
    }
});
console.log(treeCount);