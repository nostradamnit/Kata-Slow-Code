<?php
namespace SlowKata\Tests;
use SlowKata\PascalTriangle;
use PHPUnit\Framework\TestCase;

class PascalTriangleTest extends TestCase
{
    /** @test */
    public function AcceptanceTests()
    {
        // /!\ Should only pass at the end:
        // - Keep disabled
        // - Write and pass intermediate tests
        // - Enable at the end when should be passing
        // /!\ IF YOU PASS IT... YOU MIGHT HAVE GONE TOO FAST!
        $this->markTestSkipped('this test could eventually be run once a full implementation has been written');
        $pascalTriangle = new PascalTriangle();
        $this->assertEquals([1], $pascalTriangle->line(0));
        $this->assertEquals([1,1], $pascalTriangle->line(1));
        $this->assertEquals([1,2,1], $pascalTriangle->line(2));
        $this->assertEquals([1,3,3,1], $pascalTriangle->line(3));
        $this->assertEquals([1,4,6,4,1], $pascalTriangle->line(4));
        $this->assertEquals([1,5,10,10,5,1], $pascalTriangle->line(5));
        $this->assertEquals([1,6,15,20,15,6,1], $pascalTriangle->line(6));
        $this->assertEquals([1,7,21,35,35,21,7,1], $pascalTriangle->line(7));
    }
}