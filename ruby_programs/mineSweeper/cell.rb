class Cell
  
  attr_accessor :x, :y, :revealed, :display, :bee_neighbor_count, :bee
  
  def initialize(x, y)
    @x = x
    @y = y
    @revealed = false
    @display = "#"
    @bee_neighbor_count = 0
    @bee = false
  end

  def count(grid)

    if self.bee
      self.bee_neighbor_count = -1
    end

    bee_neighbor_count = 0

    (-1..1).each do |xOff|
      i = self.x + xOff
      next if i < 0 || i >= grid.length
      (-1..1).each do |yOff|
        j = self.y + yOff
        next if j < 0 || j >= grid.length

        neighbor = grid[i][j]

        if neighbor.bee
          bee_neighbor_count+=1
        end

      end
    end
    self.bee_neighbor_count = bee_neighbor_count
  end

  def reveal(grid)
    self.revealed = true
    if self.bee_neighbor_count == 0
      reveal_all_neighbors(grid)
    end
  end

  def reveal_all_neighbors(grid)
    (-1..1).each do |xOff|
      i = self.x + xOff
      next if i < 0 || i >= grid.length
      (-1..1).each do |yOff|
        j = self.y + yOff
        next if j < 0 || j >= grid.length

        neighbor = grid[i][j]

        if !neighbor.revealed
          neighbor.reveal(grid)
        end

      end
    end
  end


end
