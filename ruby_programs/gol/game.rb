class Game

  attr_accessor :grid

  def initialize()
    @grid = Array.new(15) { Array.new(15) }
    populate_grid
    print_grid
    simulate
  end


  def simulate
    system "clear"
    100.times do
      generate
      print_grid
      sleep(1)
      system "clear"
    end
  end

  def generate
    new_grid = Array.new(15) { Array.new(15) }

    @grid.each_with_index do |row, xOff|
      row.each_with_index do |col, yOff|

        neighbor_count = count_neighbors(grid, xOff, yOff)
        sum = 0
        state = @grid[xOff][yOff].state

        if state == 0 && neighbor_count == 3
          new_grid[xOff][yOff] = Cell.new(xOff, yOff, 1)
        elsif state == 1 && (neighbor_count < 2 || neighbor_count > 3)
          new_grid[xOff][yOff] = Cell.new(xOff, yOff, 0)
        else
          new_grid[xOff][yOff] = Cell.new(xOff, yOff, @grid[xOff][yOff].state)
        end

      end
    end
    @grid = new_grid
  end


  def count_neighbors(grid, x, y)
    sum = 0

    (-1..1).each do |i|
      (-1..1).each do |j|

        col = (x+i+grid.length) % grid.length
        row = (y+j+grid.length) % grid.length

        sum += grid[row][col].state
      end
    end

    sum -= grid[x][y].state
    sum
  end

  def populate_grid
    @grid.each_with_index do |row, i|
      row.each_with_index do |col, j|
        @grid[i][j] = Cell.new(i, j, rand(0..1))
      end
    end
  end

  def print_grid
    @grid.each_with_index do |row, i|
      row.each_with_index do |col, j|
        print("#{@grid[i][j].state} \t")
      end
      print("\n")
    end
  end

end

class Cell
  attr_accessor :x, :y, :state

  def initialize(x, y, state)
    @x = x
    @y = y
    @state = state
  end
end

g = Game.new()
